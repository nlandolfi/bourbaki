# generated automatically by bbk sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make roots_of_one.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

roots_of_one.tex: sheet.tex macros.tex
	bbk sheet -mode c -in sheet.tex > roots_of_one.tex

roots_of_one.pdf: ../../*.tex ../../trademark.pdf *.tex roots_of_one.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode roots_of_one.tex
	../../latexrun roots_of_one.tex
	make terms

graph.csv: sheet.tex
	bbk sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../bbk/static/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: roots_of_one.pdf
	open roots_of_one.pdf

o:
	make open

terms: roots_of_one.tex
	bbk sheet -mode ts roots_of_one.tex

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
	rm -f roots_of_one.*
	make reset
	bbk sheet -mode mk > Makefile