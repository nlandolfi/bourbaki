# generated automatically by bbk sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make normal_conditionals.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

normal_conditionals.tex: sheet.tex macros.tex
	bbk sheet -mode c -in sheet.tex > normal_conditionals.tex

normal_conditionals.pdf: ../../*.tex ../../trademark.pdf *.tex normal_conditionals.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode normal_conditionals.tex
	../../latexrun normal_conditionals.tex
	make terms

graph.csv: sheet.tex
	bbk sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../bbk/static/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: normal_conditionals.pdf
	open normal_conditionals.pdf

o:
	make open

terms: normal_conditionals.tex
	bbk sheet -mode ts normal_conditionals.tex

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
	rm -f normal_conditionals.*
	make reset
	bbk sheet -mode mk > Makefile