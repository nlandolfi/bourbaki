# generated automatically by bbk sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make set_inclusion.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

set_inclusion.tex: sheet.tex macros.tex
	bbk sheet -mode c -in sheet.tex > set_inclusion.tex

set_inclusion.pdf: ../../*.tex ../../trademark.pdf *.tex set_inclusion.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode set_inclusion.tex
	../../latexrun set_inclusion.tex
	make terms

graph.csv: sheet.tex
	bbk sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../bbk/static/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: set_inclusion.pdf
	open set_inclusion.pdf

o:
	make open

terms: set_inclusion.tex
	bbk sheet -mode ts set_inclusion.tex

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
	rm -f set_inclusion.*
	make reset
	bbk sheet -mode mk > Makefile