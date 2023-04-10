# generated automatically by bbk sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make functionals.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

functionals.tex: sheet.tex macros.tex
	bbk sheet -mode c -in sheet.tex > functionals.tex

functionals.pdf: ../../*.tex ../../trademark.pdf *.tex functionals.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode functionals.tex
	../../latexrun functionals.tex
	make terms

graph.csv: sheet.tex
	bbk sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../bbk/static/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: functionals.pdf
	open functionals.pdf

o:
	make open

terms: functionals.tex
	bbk sheet -mode ts functionals.tex

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
	rm -f functionals.*
	make reset
	bbk sheet -mode mk > Makefile