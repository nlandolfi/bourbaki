# generated automatically by bbk sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make preferences.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

preferences.tex: sheet.tex macros.tex
	bbk sheet -mode c -in sheet.tex > preferences.tex

preferences.pdf: ../../*.tex ../../trademark.pdf *.tex preferences.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode preferences.tex
	../../latexrun preferences.tex
	make terms

graph.csv: sheet.tex
	bbk sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../bbk/static/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: preferences.pdf
	open preferences.pdf

o:
	make open

terms: preferences.tex
	bbk sheet -mode ts preferences.tex

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
	rm -f preferences.*
	make reset
	bbk sheet -mode mk > Makefile