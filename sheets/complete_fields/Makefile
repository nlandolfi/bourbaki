# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make complete_fields.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

complete_fields.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > complete_fields.tex

complete_fields.pdf: ../../*.tex ../../trademark.pdf *.tex complete_fields.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode complete_fields.tex
	../../latexrun complete_fields.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: complete_fields.pdf
	open complete_fields.pdf

o:
	make open

terms: complete_fields.tex
	bbk_sheet -mode ts complete_fields.tex

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
	rm -f complete_fields.*
	make reset
	bbk_sheet -mode mk > Makefile