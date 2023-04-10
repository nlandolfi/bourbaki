# generated automatically by bbk sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make random_variable_laws.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

random_variable_laws.tex: sheet.tex macros.tex
	bbk sheet -mode c -in sheet.tex > random_variable_laws.tex

random_variable_laws.pdf: ../../*.tex ../../trademark.pdf *.tex random_variable_laws.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode random_variable_laws.tex
	../../latexrun random_variable_laws.tex
	make terms

graph.csv: sheet.tex
	bbk sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../bbk/static/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: random_variable_laws.pdf
	open random_variable_laws.pdf

o:
	make open

terms: random_variable_laws.tex
	bbk sheet -mode ts random_variable_laws.tex

spell:
	aspell -c sheet.tex

s:
	make spell

reset:
	rm -f graph.pdf
	rm -f graph.csv
	rm -f graph.graphviz

clean:
	make reset

remake:
	rm -f random_variable_laws.*
	make reset
	bbk sheet -mode mk > Makefile