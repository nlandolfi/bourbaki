# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make complete_real_inner_product_spaces.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

complete_real_inner_product_spaces.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > complete_real_inner_product_spaces.tex

complete_real_inner_product_spaces.pdf: ../../*.tex ../../trademark.pdf *.tex complete_real_inner_product_spaces.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode complete_real_inner_product_spaces.tex
	../../latexrun complete_real_inner_product_spaces.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: complete_real_inner_product_spaces.pdf
	open complete_real_inner_product_spaces.pdf

o:
	make open

terms: complete_real_inner_product_spaces.tex
	bbk_sheet -mode ts complete_real_inner_product_spaces.tex

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
	rm -f complete_real_inner_product_spaces.*
	make reset
	bbk_sheet -mode mk > Makefile