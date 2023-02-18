# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make approximators.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

approximators.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > approximators.tex

approximators.pdf: ../../*.tex ../../trademark.pdf *.tex approximators.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode approximators.tex
	../../latexrun approximators.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: approximators.pdf
	open approximators.pdf

o:
	make open

terms: approximators.tex
	bbk_sheet -mode ts approximators.tex

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
	rm -f approximators.*
	make reset
	bbk_sheet -mode mk > Makefile