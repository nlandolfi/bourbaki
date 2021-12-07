# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make surprise_datamodels.pdf

surprise_datamodels.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > surprise_datamodels.tex

surprise_datamodels.pdf: ../../*.tex ../../trademark.pdf *.tex surprise_datamodels.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode surprise_datamodels.tex
	../../latexrun surprise_datamodels.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: surprise_datamodels.pdf
	open surprise_datamodels.pdf

o:
	make open

terms: surprise_datamodels.tex
	bbk_sheet -mode ts surprise_datamodels.tex

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
	rm -f surprise_datamodels.*
	make reset
	bbk_sheet -mode mk > Makefile