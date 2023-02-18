# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make span.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

span.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > span.tex

span.pdf: ../../*.tex ../../trademark.pdf *.tex span.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode span.tex
	../../latexrun span.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: span.pdf
	open span.pdf

o:
	make open

terms: span.tex
	bbk_sheet -mode ts span.tex

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
	rm -f span.*
	make reset
	bbk_sheet -mode mk > Makefile