# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make physical_models.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

physical_models.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > physical_models.tex

physical_models.pdf: ../../*.tex ../../trademark.pdf *.tex physical_models.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode physical_models.tex
	../../latexrun physical_models.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: physical_models.pdf
	open physical_models.pdf

o:
	make open

terms: physical_models.tex
	bbk_sheet -mode ts physical_models.tex

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
	rm -f physical_models.*
	make reset
	bbk_sheet -mode mk > Makefile