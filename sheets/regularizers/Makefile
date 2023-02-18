# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make regularizers.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

regularizers.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > regularizers.tex

regularizers.pdf: ../../*.tex ../../trademark.pdf *.tex regularizers.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode regularizers.tex
	../../latexrun regularizers.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: regularizers.pdf
	open regularizers.pdf

o:
	make open

terms: regularizers.tex
	bbk_sheet -mode ts regularizers.tex

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
	rm -f regularizers.*
	make reset
	bbk_sheet -mode mk > Makefile