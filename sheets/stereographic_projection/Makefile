# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make stereographic_projection.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

stereographic_projection.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > stereographic_projection.tex

stereographic_projection.pdf: ../../*.tex ../../trademark.pdf *.tex stereographic_projection.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode stereographic_projection.tex
	../../latexrun stereographic_projection.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: stereographic_projection.pdf
	open stereographic_projection.pdf

o:
	make open

terms: stereographic_projection.tex
	bbk_sheet -mode ts stereographic_projection.tex

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
	rm -f stereographic_projection.*
	make reset
	bbk_sheet -mode mk > Makefile